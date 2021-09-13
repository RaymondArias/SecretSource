# Source Secret
Passing secrets around is difficult for me especially during development.
Instead of passing file around using a secret store then having to download the file. I wanted an app that could read a file, pull down secrets from a secret backend, SSM for now, and export values to bash.

## Usage
```bash
# build binary
go build -o sourcesecret

# Install binary
go install

# Create file with env vars mapped to secret keys
echo "export SUPER_SECRET_PASSWORD={{/test/password}}" > testfile

# Run app and secret keys will be replaced with secret value
SecretSource env -f testfile 
# App output
export SUPER_SECRET_PASSWORD='Very secure password'

# Pipe output to env vars
source <(SecretSource env -f testfile) 

# See env vars
echo $SUPER_SECRET_PASSWORD 
Very secure password


```

