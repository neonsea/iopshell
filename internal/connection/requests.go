package connection

func (c *Connection) genUbusRequest(method, path, pmethod string, message map[string]interface{}) interface{} {
    request := map[string]interface{}{
        "jsonrpc": "2.0",
        "id": c.Id,
        "method": method,
        "params": []interface{}{
            c.Key,
            path,
            pmethod,
            message,
        },
    }

    return request
}