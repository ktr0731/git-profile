# Git Profile
Dead easy git profile switcher!  

## Description
You can switch some profiles(set of username, email) with one command.  

## Equipments
- Go

## Installation
``` sh
$ go get github.com/ktr0731/git-profile
```

## Usage
Add new profile  
``` sh
$ git profile add <title> <name> <email>
```

List all profiles  
``` sh
$ git profile list
```

Use selected profile  
``` sh
$ git profile use <title>
```

Remove selected profile  
``` sh
$ git profile rm <title>
```

## License
Please see [LICENSE](./LICENSE).
