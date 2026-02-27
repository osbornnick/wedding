const AUTH_KEY = 'wedding_auth'
const CORRECT_PASSWORD = '' // change this

function isBrowser(): boolean {
  return typeof window !== 'undefined'
}

export function isAuthenticated(): boolean {
  if (!isBrowser()) return false
  return localStorage.getItem(AUTH_KEY) === 'true'
}

export function login(password: string): boolean {
  if (password === CORRECT_PASSWORD) {
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
