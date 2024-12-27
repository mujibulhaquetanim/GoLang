# CRUD - Net/Http

## close the resource

Close the response body when done because it is a resource that must be released after use to prevent memory leaks. it not closed, it will be left open and the underlying connection will not be re-used by the http client for future requests. in golang it needs to close any resources that are opened to prevent memory leaks.

## IO Reader

<details>
    <summary>In Simpler Terms</summary>

### What is `io.Reader`?

The `io.Reader` interface in Go is like a universal tool for reading data. Think of it as a standard way to grab information from different places, whether it's a file, a network connection, or even a string.

### How does it work?

- **The `Read` Method**: When you use `io.Reader`, you're essentially using its `Read` method to pull in data. This method fills up a slice of bytes (kind of like a container) with data and tells you how much data it managed to read.
  
### Simple Analogy

Imagine you have a bucket (`io.Reader`) and you want to fill it with water (data) from different sources:

- **From a tap (file)**: You can use a hose (`os.File`) to fill the bucket.
- **From a river (network)**: You can use a pipe (`net.Conn`) to get water.
- **From a water bottle (string)**: You can pour water directly into the bucket (`strings.Reader`).

No matter where the water comes from, the way you fill the bucket is the same. This is the beauty of `io.Reader`—a consistent method to read data.

### Why is it useful?

- **Flexibility**: You can read data from various sources without changing much of your code.
- **Interoperability**: It works well with many standard library functions that accept `io.Reader`.

</details>

```Go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

`p []byte`: This is a byte slice that the Read method will fill with data. It's the buffer where the data will be read into.
`n int:` This is the number of bytes actually read and written into p. It's crucial to check this value, as it might be less than the length of p.
`err error:` This indicates any error that occurred during the read operation. The most important error is io.EOF (End Of File), which signifies that there's no more data to read.

**How it Works:**

Imagine a pipe carrying data. The io.Reader is like the end of that pipe where you can take data from.

You provide a buffer (`p []byte`) to the Read method.
The Read method attempts to fill that buffer with data from the underlying source (e.g., a file, network connection, memory buffer).
Read returns the number of bytes read (n) and any error that occurred.

**Key Concepts:**

*Streaming:* The io.Reader interface enables streaming. You don't need to load the entire data source into memory at once. You read it in chunks, which is very efficient for large files or network streams.
*Abstraction*: io.Reader abstracts away the underlying data source. Your code that uses an io.Reader doesn't need to know if it's reading from a file, a network connection, or an in-memory buffer.

**Several types in the Go standard library implement the io.Reader interface:**

`os.File:` Represents a file on the file system.
`bytes.Reader:` Reads from an in-memory byte slice. This is what we use when sending JSON data in HTTP requests (as shown in the previous examples).
`strings.Reader:` Reads from an in-memory string.
`net.Conn:` Represents a network connection.
`bufio.Reader:` Provides buffered reading, which can improve performance by reducing the number of system calls.

## Decode/Encode

Decode the response body into a struct type Todo using json.NewDecoder and json.NewDecoder.Decode. json.NewDecoder returns a new decoder that reads from res.Body. json.NewDecoder.Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by &todo. it is preferred to use json.NewDecoder over json.Unmarshal when reading from an io.Reader because it is more efficient to read from an io.Reader than to read from a byte slice. json.NewDecoder also allows you to read from a stream of JSON-encoded values, whereas json.Unmarshal expects a single JSON-encoded value. it is also more efficient to use json.NewDecoder when you don't know the size of the JSON-encoded value you are reading because it reads the JSON-encoded value in chunks and decodes it incrementally, whereas json.Unmarshal reads the entire JSON-encoded value into memory before decoding it. json.NewDecoder also allows you to read from an io.Reader that is not a byte slice, such as a file or a network connection.

- json.NewDecoder is more suited for reading large or streaming JSON data, while json.Unmarshal is simpler and convenient for handling smaller, in-memory JSON data.

## Reading data using io and its ReadAll method

ReadAll reads from the response body until an error or EOF and returns the data it read. A successful call returns err == nil, not err == EOF. Because ReadAll is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.

## POST Request handling

- The whole flow is as follows: struct -> JSON -> byte array -> string -> io.Reader -> POST request -> response -> byte array -> string

convert struct to JSON -> byte array using json.Marsal -> json.Marshal returns byte array and error -> check error -> convert byte array to string using string() function -> convert string to io.Reader using strings.NewReader  -> use io.Reader to make POST request using http.Post -> http.Post returns response and error -> check error -> check response status code -> read bytes from response body using io.ReadAll -> io.ReadAll returns byte array and error -> check error -> convert byte array to string using string() function -> print the string. The same can be done using json.NewDecoder and json.NewEncoder. The difference is json.NewDecoder and json.NewEncoder works with io.Reader and io.Writer interfaces respectively. json.Marshal and json.Unmarshal works with byte array. json.Marshal returns byte array and error. json.Unmarshal returns error. json.NewDecoder and json.NewEncoder returns pointer to Decoder and Encoder respectively. Decoder and Encoder works with io.Reader and io.Writer interfaces respectively.
