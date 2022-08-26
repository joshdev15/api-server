const http = require('http')
const port = 8080

/** Server */
const server = http.createServer(async (req, res) => {
	try {
		const path = req.url
		let value = {}
		let code = 200

		/** Pseudo Router */
		const options ={
			'/' : () => {
  			value = {'title': 'Simple API Rest Nodejs example', 'success': 'Hola server en Node.js!', 'statusCode': 200}
  			code = 200
			},
			'/error': () => {
  	  	value = {'title': 'Simple API Rest Nodejs example', 'error': 'Esto es un error!', 'statusCode': 500}
  			code = 500
			}
		}

		if (options[path]) {
			options[path]()
			res.writeHead(code, {'Content-Type': 'application/json'})
			res.end(JSON.stringify(value))
		}
	} catch (error) {
		console.log(error)
	}
})

/** Server Start */
server.listen(port, () => {
	console.log(`Your server is running! - http://localhost:${port}`)
})

/** Made in Node 16.16.0 */
