# Microservice

## Setup

As the environment was dockerized, you need to have docker and docker-compose tools in order to be able to run the microserve.

Command to run app:
```
docker-compose up
```


## Explanation

These are the fields for the `product` model.

- Id (uint)
- Name (string) **mandatory**
- Supplier Id (uint) 
- Category Id (uint)
- Units In Stock (uint) **mandatory**
- Unit Price (float64) **mandatory**
- Discontinued (bool)

As you can see there are three **mandotory** fields, that means if you dont send them within the request in which are required the response will be an error.

The following block are the available endpoints of the API.

```ruby
                    POST /api/products                          
                    GET /api/products/{id}             
                    PUT /api/products/{id}                
                    DELETE /api/products/{id}                
``` 
You must send the request as the following example:

``
  localhost:8080/api/products/1
``

Finally, as long as you are in the root app directory you can run the tests with the following command:

``
  golang test ./... -v -cover
``
