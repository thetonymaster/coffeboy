# Coffeboy

##Setup
First you should have a file structure like this: ````$GOPATH/src/github.com/crowdint```` then clone this repo inside the ````crowdint```` folder.

Then get the godeps package ````go get github.com/tools/godep````.

Once downloaded, you should add the the bin folder to your PATH in your .bashrc file or equivalent:
````export PATH=$PATH:$GOPATH/bin````, and then reload or open a new terminal.

Now you will be able to run ````godep restore```` to download all the depndencies.

##Examples

###Orders
Create

    curl -i -X POST http://coffeboy.herokuapp.com/order --data-binary '{  
       "id":"R9999",
       "user_id":1,
       "created_at":"",
       "updated_at":"",
       "completed_at":"",
       "email":"",
       "total_quantity":"",
       "line_items":[  
          {  
             "variant_id":"1",
            "quantity":10
          },
          {  
             "variant_id":"2",
             "quantity":20
          }
       ]
    }'

Get

    curl -i http://coffeboy.herokuapp.com/order/R9999
    
Update

    curl -i -X PUT http://coffeboy.herokuapp.com/order/R9999 --data-binary '{  
       "id":"R9999",
       "user_id":1,
       "created_at":"",
       "updated_at":"",
       "completed_at":"",
       "email":"someemail@email.com",
       "total_quantity":"",
       "line_items":[  
          {  
             "variant_id":"1",
             "quantity":10
          },
          {  
             "variant_id":"2",
             "quantity":20
          }
       ]
    }'
    
Delete

    curl -i -X DELETE http://coffeboy.herokuapp.com/order/R9999
