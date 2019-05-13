# Dodo Util
A tiny library for interfacing with a Dodo document store in Go!

## Using Dodo Util
This can be installed just as any other Go dependency: ```go get github.com/danielwoodsdeveloper/dodo-util```. Import it to your app: ```import ("danielwoodsdeveloper/dodo-util")```.

This gives you access to a number of easy-to-use functions:
- ```Store(doc map[string]interface{})```
- ```Get(id string)```
- ```Modify(id string, doc map[string]interface{})```
- ```GetAll()```
- ```Delete(id string)```
- ```DeleteAll()```

To set up the Dodo connection in the first place:

```
st := dodo.Settings{"http://localhost:6060", "", "", "", ""}
d, err := dodo.NewDodoConnection(st)
if err != nil {
    panic(err)
}
```