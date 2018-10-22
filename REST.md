

__Warning: Do not use any path marked PRIVATE, they will begin to require credentials in the near future__

###GET /servers/
Get the current registered servers
```
HTTP/1.1 200 OK
Date: Mon, 22 Oct 2018 15:31:18 GMT
{
	"server_name":{
		"Addr":(string),
		"Iaddr":{
			"Ip":(string),
			"Gateway":(string),
			"Subnet":(int)
		},
		"Nodes":(int),
		"Max":(int),
		"Id":(int),
		"Iface":(string),
		"Switches":[
			{
				"Addr":(string),
				"Iface":(string),
				"Brand":(int),
				"Id":(int)
			}
		]
	},
	"server2_name":{...}...
}
```


###PUT /servers/{name}
####PRIVATE
Register and add a new server to be 
controlled by the instance
```
{
	"Addr":(string),
	"Iaddr":{
		"Ip":(string),
		"Gateway":(string),
		"Subnet":(int)
	},
	"Nodes":(int),
	"Max":(int),
	"Id":-1,
	"Iface":(string),
	"Switches":[
		{
			"Addr":(string),
			"Iface":(string),
			"Brand":(int),
			"Id":(int)
		}
	]
}
```

```
HTTP/1.1 200 OK
Date: Mon, 22 Oct 2018 15:31:18 GMT
<server id>
```

###GET /servers/{id}
Get a server by id
```
HTTP/1.1 200 OK
Date: Mon, 22 Oct 2018 15:31:18 GMT
{
	"Addr":(string),
	"Iaddr":{
		"Ip":(string),
		"Gateway":(string),
		"Subnet":(int)
	},
	"Nodes":(int),
	"Max":(int),
	"Id":(int),
	"Iface":(string),
	"Switches":[
		{
			"Addr":(string),
			"Iface":(string),
			"Brand":(int),
			"Id":(int)
		}
	]
}
```

__PRIVATE__
DELETE /servers/{id}

HTTP/1.1 200 OK
Date: Mon, 22 Oct 2018 15:31:18 GMT
Success


UPDATE /servers/{id}
{
	"Addr":(string),
	"Iaddr":{
		"Ip":(string),
		"Gateway":(string),
		"Subnet":(int)
	},
	"Nodes":(int),
	"Max":(int),
	"Id":(int),
	"Iface":(string),
	"Switches":[
		{
			"Addr":(string),
			"Iface":(string),
			"Brand":(int),
			"Id":(int)
		}
	]
}

HTTP/1.1 200 OK
Date: Mon, 22 Oct 2018 15:31:18 GMT
Success


GET /testnets/

HTTP/1.1 200 OK
Date: Mon, 22 Oct 2018 15:31:18 GMT
[
	{
		"Id":(int),
		"Blockchain":(string),
		"Nodes":(int),
		"Image":(string)
	},...

]


POST /testnets/
{
	"Servers":[(int),(int)...],
	"Blockchain":(string),
	"Nodes":(int),
	"Image":(string)
}

HTTP/1.1 200 OK
Date: Mon, 22 Oct 2018 15:31:18 GMT
Success


GET /testnets/{id}

HTTP/1.1 200 OK
Date: Mon, 22 Oct 2018 15:31:18 GMT
{
	"Id":(int),
	"Blockchain":(string),
	"Nodes":(int),
	"Image":(string)
}

__PRIVATE__
GET /testnets/{id}/nodes/

HTTP/1.1 200 OK
Date: Mon, 22 Oct 2018 15:31:18 GMT
[
	"Id":(int),
	"TestNetId":(int),
	"Server":(int),
	"LocalId":(int),
	"Ip":(string)
]

