# Golang client library for MultiChain blockchain

## Usage

```

  chain := flag.String("chain", "", "is the name of the chain")
  host := flag.String("host", "localhost", "is a string for the hostname")
  port := flag.String("port", "80", "is a string for the host port")
  username := flag.String("username", "multichainrpc", "is a string for the username")
  password := flag.String("password", "12345678", "is a string for the password")

  flag.Parse()

  client := NewClient(
      *chain,
      *host,
      *port,
      *username,
      *password,
  )
    
  obj, err := client.GetInfo()
  if err != nil {
      panic(err)
  }
  
  fmt.Println(obj)

```
