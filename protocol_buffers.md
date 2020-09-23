# Introduction

###### Protocol Buffer is a kind of  encoded structured data similar to JSON with advantages:
- smaller
- faster
- simpler

###### Use Protocol Buffers to exchange data bring great performance.

# Protocol Buffers Interface Definition Language
###### Protocol buffers are saved as `*.proto` file.

```
syntax = "proto3";

message SearchRequest {
    string query = 1;
    int32 page_number = 2;
    optional int32 items_per_page = 3 [default=10];

    reserved 4, 5;
    reserved "foo", "bar";

    enum SearchType {
        UNIVERSAL = 0;
        WEB = 1;
        IMAGE = 2;
        NEWS = 3;
        VIDEO = 4;
    }
    optional SearchType type = 1;

    enum Status {
        option allow_alias = true;
        INACTIVE = 0;
        ACTIVE = 1;
        RUNNING = 1;
    }
}
```

###### The first line declares the syntax version proto file, default is version 2. The first line of proto file must be non-empty, non-comment

Field number use to identify fields in binary message format, should not be changed once your message type is in use. Number from 1 to 15 take 1 byte to encode (should use in frequent message), number from 16 to 2047 take 2 bytes to encode.

###### Field rules: singular and repeated.
###### Reserved field
###### Field can be `optional` or `required`
###### You can define an Enumuration Type for pre-defined value list. We can alias value in Enum by using `option allow_alias = true;`