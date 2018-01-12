# Redirect

this is a redirect server for many use.

## Quick start

build executable binary file

```bash
go get -d 
go build 
```

execute 

```bash
./redirect --conf config.yaml
```

a simple config file like following:

```yaml
static:
  data:
    xx: yy
    zz: pp
    var-service: service
  template:
    _g: (refreshInterval:(display:Off,pause:!f,value:0),time:(from:'{{ $.from }}',mode:absolute,to:'{{ $.to }}'))
    _a: (columns:!(_source),index:{{ $.service }},interval:auto,query:(language:lucene,query:'{{ $.query }}'),sort:!('@timestamp',desc))
  time:
    keys:
      - from
      - to
    from:
      format: timestamp
      unit: ms
    to:
      format: 2006-01-02T15:04:05.000Z
      unit: ms
      tz: Asia/Shanghai
    skip:
      - now.*
destination:
  all: http://kibana.vinkdong.com/app/kibana#/discover
```

- In <b>static.data</b> transfer key xx-->yy,zz-->pp,var-service-->service,such as source link is `localhost?xx=100&&zz=99` will transfer to `http://kibana.vinkdong.com/app/kibana#/discover?yy=100&&pp=99`

- In <b>static.template</b> it will add the key to the parameter and execute the template values

- In <b>static.time</b> tail program which key is time key, and transfer to individual time format, the all-time format in standard goland have supported, BTW you can ignore some particular time value by regex

- In <b>destination.all</b> default destination