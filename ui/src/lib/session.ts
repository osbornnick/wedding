import { useSession } from '@tanstack/react-start/server'

export interface User {
  ip: string
  x_forwarded_for?: string
  name?: string
  hasRSVPed?: boolean
}

export interface AuthState {
  isAuthenticated: boolean
  user: User | null
}

export function useAppSession() {
  return useSession<AuthState>({
    // Session configuration
    name: 'wedding-session',
    password: process.env.SESSION_SECRET!, // At least 32 characters
    // Optional: customize cookie settings
    cookie: {
      secure: process.env.NODE_ENV === 'production',
      sameSite: 'lax',
      httpOnly: true,
    },
  })
}
