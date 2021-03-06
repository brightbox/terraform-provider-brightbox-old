package brightbox

import (
	"fmt"
	"log"
	"regexp"

	"github.com/brightbox/gobrightbox"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceBrightboxDatabaseType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBrightboxDatabaseTypeRead,

		Schema: map[string]*schema.Schema{

			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"disk_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"ram": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceBrightboxDatabaseTypeRead(
	d *schema.ResourceData,
	meta interface{},
) error {
	client := meta.(*CompositeClient).ApiClient

	log.Printf("[DEBUG] DatabaseType data read called. Retrieving database type list")

	databaseTypes, err := client.DatabaseServerTypes()
	if err != nil {
		return fmt.Errorf("Error retrieving database type list: %s", err)
	}

	databaseType, err := findDatabaseTypeByFilter(databaseTypes, d)

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Single DatabaseType found: %s", databaseType.Id)
	return dataSourceBrightboxDatabaseTypesAttributes(d, databaseType)
}

func dataSourceBrightboxDatabaseTypesAttributes(
	d *schema.ResourceData,
	databaseType *brightbox.DatabaseServerType,
) error {
	log.Printf("[DEBUG] databaseType details: %#v", databaseType)

	d.SetId(databaseType.Id)
	d.Set("name", databaseType.Name)
	d.Set("description", databaseType.Description)
	d.Set("disk_size", databaseType.DiskSize)
	d.Set("ram", databaseType.RAM)

	return nil
}

func findDatabaseTypeByFilter(
	databaseTypes []brightbox.DatabaseServerType,
	d *schema.ResourceData,
) (*brightbox.DatabaseServerType, error) {
	nameRe, err := regexp.Compile(d.Get("name").(string))
	if err != nil {
		return nil, err
	}

	descRe, err := regexp.Compile(d.Get("description").(string))
	if err != nil {
		return nil, err
	}

	var results []brightbox.DatabaseServerType
	for _, databaseType := range databaseTypes {
		if databaseTypeMatch(&databaseType, d, nameRe, descRe) {
			results = append(results, databaseType)
		}
	}
	if len(results) == 1 {
		return &results[0], nil
	} else if len(results) > 1 {
		return nil, fmt.Errorf("Your query returned more than one result (found %d entries). Please try a more "+
			"specific search criteria.", len(results))
	} else {
		return nil, fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}
}

//Match on the search filter - if the elements exist
func databaseTypeMatch(
	databaseType *brightbox.DatabaseServerType,
	d *schema.ResourceData,
	nameRe *regexp.Regexp,
	descRe *regexp.Regexp,
) bool {
	_, ok := d.GetOk("name")
	if ok && !nameRe.MatchString(databaseType.Name) {
		return false
	}
	_, ok = d.GetOk("description")
	if ok && !descRe.MatchString(databaseType.Description) {
		return false
	}
	return true
}
