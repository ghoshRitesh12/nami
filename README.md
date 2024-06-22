```
                       _
                      (_)
 _ __   __ _ _ __ ___  _
| '_ \ / _` | '_ ` _ \| |
| | | | (_| | | | | | | |
|_| |_|\__,_|_| |_| |_|_|
```

An attempt for a File Based Route Generator in Go. <br/>
It's named after a fictional character in the One Piece franchise, who is the **navigator** of the Staw Hat Pirates.

## [Conventions](#convention)

Nami follows a mixture of file based routing conventions that [NextJS](https://nextjs.org/) and [NuxtJS](https://nuxt.com/) uses.

<!-- - -->

### [Main directory](#main-directory)

Main directory refers to the root directory from which routes may originate. The name of the main directory should be **routes**. If this directory is not found, then an error is returned.

### [Possible Route Handler File Names](#possible-route-handler-file-names)

```sh
"route.get.go" #or
"route.put.go" #or
"route.post.go" #or
"route.head.go" #or
"route.patch.go" #or
"route.delete.go"
```

Any file present within the **routes** directory that doesn't match the above possible file names would be excluded from route mapping

### [Generated Route File](#generated-route-file)

After traversing through the [main directory](#main-directory) and skipping files that doesn't follow convention, a `main.gen.go` file is automatically generated that contains all the route definitions along with their handler and verb.

> TODO:
>
> Figure out how to use `go generate` to generate prebuild route mappings and also perhaps do some stuff to integrate the generate route mappings with a web framework.
>
> This should be fun
