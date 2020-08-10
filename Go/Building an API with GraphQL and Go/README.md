https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091

https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356

## Set golang env to be used globally

```

mkdir -p ~/go/{bin,pkg,src}

echo 'export GOPATH="\$HOME/go"' >> ~/.bashrc

echo 'export PATH="$PATH:\${GOPATH//://bin:}/bin"' >> ~/.bashrc

```
