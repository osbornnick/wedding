import { createServerFn } from '@tanstack/react-start'
import { API_HOST } from './constants'

export type Guest = {
  name: string
  id: string
  invitationId: string
  aliases: string[]
  createdAt: string
}

export const findGuestsFn = createServerFn()
  .inputValidator((data: { name: string }) => data)
  .handler(async ({ data }) => {
    console.log('Finding guests for name:', data.name)
    const guests = await findGuests(data.name)
    console.log('Found guests:', guests)
    return guests
  })

async function findGuests(name: string): Promise<Guest[] | null> {
  const response = await fetch(
    `${API_HOST}/api/guests?name=${encodeURIComponent(name)}`,
  )
  if (response.ok) {
    return response.json()
  }
  return null
}