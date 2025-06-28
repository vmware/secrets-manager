
## THINGS TO DO FOR THE BIG SUMMER CLEANUP

* Build and test docs locally and configure CloudFlare to point to the ./public folder instead.
* Decommission DockerHub, use GCR instead.
* Automate container creation.
* Remove k8s manifest generation; `helm` is the only way to deploy to a cluser.
* Remove Carvel support and other custom extensions.
* Create an `./experimental` folder and move all experimental things over (*we 
  can start working on them depending on community support and project direction*) 
* Deploy SPIRE from helm charts.
* Deploy VSecM locally from helm charts first and verify things work.
* Update doc versioning; likely remove references to <v.40 docs; better to 
  start fresh.
* Remove FIPS config option; FIPS will be supported by default and always.
* Update examples and use cases in docs, since some of them are moved to the
  `experimental` folder and will need re-testing, re-iterating.
* Remove SSF badge; we can plan for it and add it back once we have >80% 
  test coverage and fuzzy testing.
* 