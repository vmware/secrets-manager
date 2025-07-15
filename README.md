## **VMware Secrets Manager** *for cloud-native apps*

![VSecM Logo](https://github.com/vmware/secrets-manager/assets/1041224/885c11ac-7269-4344-a376-0d0a0fb082a7)

## VMware Secret Manager v2.0

### 🚀 Major Architectural Shift: SPIKE-Powered Backend

VSecM v2 represents a fundamental reimagining of VSecM architecture. We're 
transitioning from a standalone secrets store to a comprehensive lifecycle 
manager and orchestration layer for SPIFFE-first secrets management.

## Status: Experimental

🧪 **v2.0.0 is currently experimental**.

We're actively developing and stabilizing the SPIKE integration.

**Expect frequent updates and potential breaking changes until we reach v2.1.0**.

**For production use, we recommend staying on v0.x until v2 stabilizes**.

### What's New in v2

**SPIKE Integration**: VSecM now uses [SPIKE](https://spike.ist) as its backend 
storage engine. SPIKE is a minimal, SPIFFE-native secrets store providing:

* Better durability and consistency guarantees
* Composable architecture
* True SPIFFE-first design

**New Architecture**:

- **VSecM Safe** routes all operations to SPIKE backend (Phase 1: **WORK IN PROGRESS**)
- **File-based storage** is deprecated in favor of SPIKE's SQLite store
- **API compatibility** maintained for smooth migration

**Container Registry Change**:

We are moving from **Docker Hub** and using 
[**Github Container Registry**](https://docs.github.com/en/packages) for 
better streamlining and automation of container image creation.

**Helm-First**:

We are phasing out custom Kubernetes manifests; all deployment will be through
Helm Charts.

**Removal of Experimental Features**:

We are starting this version as lean as possible; we will remove all experimental
features and existing use cases and examples. Since this is a major architecture
change, we will have to update existing examples and create new examples for
the newer use cases. So, we decided to start from scratch altogether.

### Why This Change?

VSecM and SPIKE serve complementary roles in the SPIFFE ecosystem:
- **SPIKE**: Minimal, secure, SPIFFE-native storage layer
- **VSecM**: Full lifecycle management, orchestration, UI, and fleet capabilities

By combining them, we remove duplication and create a more powerful, modular system.

### Migration from v0.x

⚠️ **Breaking Changes**: This is a complete architectural overhaul. 
While we will do our best to let APIs remain compatible, the underlying storage 
mechanism will change fundamentally. The best way to migrate would be to 
export your secrets, and create an automation script to recreate them in
the new setup.

### Roadmap

- [ ] Phase 1: SPIKE as a pluggable backend
- [ ] Phase 2: Deprecate file-based storage
- [ ] Phase 3: WebCrypto-secured UI (Tornjak-like)
- [ ] Phase 4: Full lifecycle management & fleet orchestration

<!-- TODO: add links to CoC, contributing guidelines, etc. -->
