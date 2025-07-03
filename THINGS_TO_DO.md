
## THINGS TO DO FOR THE BIG SUMMER CLEANUP

* Build and test docs locally and configure CloudFlare to point to the ./public folder instead.
* Decommission DockerHub, use GCR instead.
* Automate container creation.
* Remove k8s manifest generation; `helm` is the only way to deploy to a cluster.
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
* Remove the SSF badge; we can plan for it and add it back once we have >80% 
  test coverage and fuzzy testing. also create issues for it.
* Do we really need VSecM Keystone?
* deprecate `age` support, default to FIPS algorithms only.
* `vsecmsystem` -> `vsecm` (we are not doing any system operations wrt k8s context)
* Make SPIRE helm charts manage vsecm clusterspiffeids.
* have a special secret kind for SPIKE; a way to transform it back and forth.
* Read and update all documentation since a lot will change.
* remove spire from subcharts; we will install it from upstream directly.
* remove keystone and scout
* remove crds; SPIRE chart already includes them.
* remove most of the examples as depending on the project direction they may need to be revised.
* mac-tunnel: changed; there is a better way to do it via port forwarding. -- spike does that already.
* simplify the matrix of dockerfiles; too many combination there.