# Echo Practise With MongoDb


## Mongo DB
- Document Oriented database 
- Non relational Database
- Document databases store all information for a given object in a single instance in the database and every stored object can be different from every other.
- Documents in a document store are roughly equivalent to the programming concept of an object.They are not requierd to adhere to standard schema
- Are addresssed in the database via a unique key that represents that document.
- Document stores use the metadata in the document to classify the content
- collection is called as tables
- document is called as rows
- organisation is called as Database

### ObjectID
  - mongodb with have a objectId
  - made up of 12 bytes
  - contains information about timestamp,machine id, process id, counter

### BSON Data
- The Go driver provides four main types for working with BSON data:
  - D: An ordered representation of a BSON document (slice)
      - bson.D{{"hello","world"},{"foo","bar"}}
  - M: An unordered representation of a BSON document (map)
      - bson.M{"hello":"world","foo":"bar"}
  - A: An ordered representation of a BSON array
      - bson.A{"Eric","Tests"}
  - E: A single element inside a D type


## Sample API
### POST
```json
{
  "product_name":"one",
  "vendor":"apple",
  "email":"test@gmail.com",
  "website":"https://apple.com",
  "country":"US",
  "default_device_ip":"192.168.9.8"
}
```