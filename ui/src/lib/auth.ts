const AUTH_KEY = 'wedding_auth'

function isBrowser(): boolean {
  return typeof window !== 'undefined'
}

export function isAuthenticated(): boolean {
  if (!isBrowser()) return false
  return localStorage.getItem(AUTH_KEY) === 'true'
}

export async function login(password: string): Promise<boolean> {
  const response = await fetch('http://localhost:8081/api/users/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username: 'user', password: password }),
  })
  console.log(response)
  if (response.ok) {
    if (isBrowser()) {
      localStorage.setItem(AUTH_KEY, 'true')
    }
    return true
  }
  return false
}

export function logout(): void {
  if (!isBrowser()) return
  localStorage.removeItem(AUTH_KEY)
}
