# CRUD - Net/Http

## close the resource

Close the response body when done because it is a resource that must be released after use to prevent memory leaks. it not closed, it will be left open and the underlying connection will not be re-used by the http client for future requests. in golang it needs to close any resources that are opened to prevent memory leaks.

## Decode/Encode

Decode the response body into a struct type Todo using json.NewDecoder and json.NewDecoder.Decode. json.NewDecoder returns a new decoder that reads from res.Body. json.NewDecoder.Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by &todo. it is preferred to use json.NewDecoder over json.Unmarshal when reading from an io.Reader because it is more efficient to read from an io.Reader than to read from a byte slice. json.NewDecoder also allows you to read from a stream of JSON-encoded values, whereas json.Unmarshal expects a single JSON-encoded value. it is also more efficient to use json.NewDecoder when you don't know the size of the JSON-encoded value you are reading because it reads the JSON-encoded value in chunks and decodes it incrementally, whereas json.Unmarshal reads the entire JSON-encoded value into memory before decoding it. json.NewDecoder also allows you to read from an io.Reader that is not a byte slice, such as a file or a network connection.
- json.NewDecoder is more suited for reading large or streaming JSON data, while json.Unmarshal is simpler and convenient for handling smaller, in-memory JSON data.

## Reading data using io and its ReadAll method

ReadAll reads from the response body until an error or EOF and returns the data it read. A successful call returns err == nil, not err == EOF. Because ReadAll is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.
