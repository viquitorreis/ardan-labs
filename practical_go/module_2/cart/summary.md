Slices não são arrays. Slices são como uma **view** do array por baixo dos panos.

Struct: 

```
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```