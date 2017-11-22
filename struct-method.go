type T struct {
	name string 
	value int
}
	
func (t T) String() string {
	return fmt.Sprintf("name %s, values %d", t.name, t.value)
}

func (t *T) Update(i int) string {
	t.value = i
}