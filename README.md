# HMAC-SHA256 vs Ed25519 vs ES256

This is a repository that benchmarks the performance of signing/authenticating messages using either:

- HMAC with SHA-256
- Ed25519
- secp256r1 with SHA-256 (ES256)

The purpose of benchmarking is to compare the performance of HMAC as opposed to elliptic curve signature for message authentication.

The conclusion: HMAC far outperform elliptic curve digital signatures.

Here are the raw benchmark results when run on my 2020 M1 MacBook Pro:

```
BenchmarkEd25519-8                 14067             83480 ns/op
BenchmarkHmacSha256-8            1905396               626.3 ns/op
BenchmarkSecp256r1-8               13813             84879 ns/op
```

## Motivation for even caring about the performance

In an application that I am working on, messages will need to be relayed from node to node. Bad actors will need to be booted off the network.

Among all the definitions of a bad actor, one such being that the actor is intentionally tampering with data. Actors that are caught tampering must be removed from the network.

With that said, how do we prove that the bad actor did indeed tamper with the message?

Plenty of solutions available, each with their pros and cons. Symmetric vs asymmetric cryptography are two solutions to explore. Given the nature of the problem, performance is a major decision that needs to be factored in.
