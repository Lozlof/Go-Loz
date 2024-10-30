## Go Slice   
https://go.dev/blog/slices-intro    
#### A slice is a reference type in Go that represents a segment of an underlying array.   
#### Slices are more flexible than arrays because their size can grow or shrink dynamically.    
**Dynamic resizing:** Although slices have a fixed size once initialized, you can create new slices with larger capacities using the built-in append() function.       
**Underlying array:** A slice always points to an underlying array. Modifying the slice will modify the original array.        
**Properties:**      
- ``len(slice)``: The length of the slice.              
- ``cap(slice)``: The capacity of the slice, which is the maximum length the slice can grow without allocating more memory.        
#### From Go:     
A slice is like an array, except that its size changes dynamically as you add and remove items. The slice is one of Go's most useful types.