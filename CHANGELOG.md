## 1.0.6 (Unreleased)
## 1.0.5 (August 01, 2018)
- Terraform provider official release

## 1.0.4 (July 2, 2018)

- Rewrite structures to standard format
- Refactor Load Balancer
- Refactor Cloud Ip
- Refactor Firewall Policy
- Refactor Firewall Rule
- Refactor Server Group
- Add sensitive marker to container client key
- Remove redundant returns in brightbox server
- Refactor Container code
- Refactor Validate Functions
- Add sensitive marker the provider password entry
- Enforce healthcheck max items in schema
- Remove redundant test in cloudip
- Refactor image test to latest ubuntu image
- Use map_from_string_set function
- Use make rather than append to create string slice
- Improve name of curried function
- Tidy up code
- Allow tests to run with 1.10
- Update vendor

## 1.0.3 (March 13, 2018)

- Make database type check region independent

## 1.0.2 (March 9, 2018)

- On correct UUID
- Access un/pw github credentials properly in Jenkinsfile

## 1.0.1 (March 9, 2018)

- Switch github credentials
- Remove incorrect database fields

## 1.0.0 (March 8, 2018)

- Update README
- Update weblayer example
- Update website docs
- Add Database Type Support
- Add tagrelease shell script
- Add Database Type support

## 0.1.4 (March 8, 2018)

- Remove Jenkinsfile debugging
- Update github-token
- Set debug mode on goreleaser
- Don't do update of get-tools
- Refactor release process to makefile
- Add network capacity to docker container
- Add release process to Jenkinsfile
- Fetch tags
- Add goreleaser snapshot build for non-master branches
- Improve Orbit URL manipulation
- Fix Credentials
- Copy back junit report
- Run test from same directory
- Only create directories to dirname of git remote
- Quote shell script
- Move workspace to Go layout
- Initial Jenkinsfile
- Switch to dep vendoring
- Update index.html.markdown
- Update gitignore
- Bump version
- Ensure static binaries

## 0.1.3 (January 15, 2018)

- Ensure static binaries

## 0.1.2 (September 6, 2017)

- Add plugin installer script
- Allow userdata update
- Add encoded userdata option - to support templates
- Document resources

## 0.1.1 (August 22, 2017)

- Add import and deletion support for more resources

## 0.1.0 (August 21, 2017)

- Add brightbox image data source support
- Update provider to new breakout format
- Make sure terraform refresh removes deleted servers
- Add import support for servers
