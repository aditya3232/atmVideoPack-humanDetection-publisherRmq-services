
# API Spec Publisher AtmVideoPack Human Detection

## 1.1 Human Detection

### 1.1.1 POST :: Human Detection

Request :
- Method : POST
- Endpoint : `localhost:3333/publisher/atmvideopack/v1/humandetection/create`
- Header :
    - Content-Type : application/x-www-form-urlencoded
    - x-api-key : required
- Body (form-data: x-www-form-urlencoded) :
    - tid : string, required
    - date_time : string, required
    - person : string, required
    - file : file, required
- Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer",
    },
        "data":{
            "tid": "string",
            "date_time": "string",
            "person": "string",  
            "converted_file": "string"
        }
 }
```






