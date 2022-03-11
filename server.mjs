/* eslint-disable @typescript-eslint/no-misused-promises */
// @ts-check

import dotenv from 'dotenv'
import { readFileSync } from 'fs'
import { createServer } from 'https'
import next from 'next'

dotenv.config()

const dev = process.env.NODE_ENV !== 'production'
const hostname = process.env.HOSTNAME ?? 'localhost'
const port = process.env.PORT ? parseInt(process.env.PORT) : 3000

const app = next({
  dev,
  hostname,
  port,
})

// mkcert certificate
const server = createServer(
  {
    cert: readFileSync('./localhost.pem'),
    key: readFileSync('./localhost-key.pem'),
  },
  app.getRequestHandler()
)

// @ts-ignore
await app.prepare()

server.listen(port, () => {
  console.log(`> Ready on https://${hostname}:${port}`)
})
