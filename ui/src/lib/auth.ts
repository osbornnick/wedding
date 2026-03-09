import { createServerFn } from '@tanstack/react-start'
import { useAppSession } from './session'
import { redirect } from '@tanstack/react-router'
import { getRequestIP } from '@tanstack/react-start/server'
import { API_HOST } from './constants'

export const loginFn = createServerFn({ method: 'POST' })
  .inputValidator((data: { password: string }) => data)
  .handler(async ({ data }) => {
    const success = await login(data.password)
    if (success) {
      const session = await useAppSession()
      await session.update({
        user: {
          ip: getRequestIP() || 'unknown',
          x_forwarded_for: getRequestIP({ xForwardedFor: true }),
        },
        isAuthenticated: true,
      })
    }
    return success
  })

async function login(password: string): Promise<boolean> {
  const response = await fetch(`${API_HOST}/api/users/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username: 'user', password: password }),
  })
  console.log(response)
  if (response.ok) {
    return true
  }
  return false
}

export const logoutFn = createServerFn({ method: 'POST' }).handler(async () => {
  const session = await useAppSession()
  await session.clear()
  throw redirect({ to: '/login' })
})
