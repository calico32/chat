import md5 from 'md5'
import { ClientSocketMessage } from '../server/types'

const cleanEmail = (email: string): string => {
  return email.toLowerCase().trim()
}

export const gravatarUrl = (email: string): string => {
  const hash = md5(cleanEmail(email))
  return `https://www.gravatar.com/avatar/${hash}`
}

export const hashSecret = async (username: string, secret: string): Promise<string> => {
  const plain = new TextEncoder().encode(`${username}:${secret}`)
  const hash = await crypto.subtle.digest('SHA-512', plain)
  return Array.from(new Uint8Array(hash))
    .map((b) => b.toString(16).padStart(2, '0'))
    .join('')
}

export const sendMessage = <T extends ClientSocketMessage>(
  ws: WebSocket,
  type: T['type'],
  data: T['data']
): void => {
  ws.send(JSON.stringify({ type, data }))
}
