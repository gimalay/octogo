# Octogo

It's a Proof of Concept / Research project. 

The goal of the project is to create a mobile application with native iOS and Android UI's where the majority of the logic is shared and implemented in Go. 

# Techinical details 
- Persitance layer is implement via boltdb embeded key-value store
- Tiny Command/Query API between UI and application logic 
- Static typing of ViewModels and Commands implemented with Protobuf codegeneration for Go, Swift and Java

# Swift UI 
Swift UI connector
https://github.com/gimalay/octogo.swift/blob/master/README.md
