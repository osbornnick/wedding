import { createFileRoute, useNavigate } from '@tanstack/react-router'
import { useState } from 'react'
import { login, isAuthenticated } from '../lib/auth'

export const Route = createFileRoute('/login')({
  component: LoginPage,
})

function LoginPage() {
  const navigate = useNavigate()
  const [password, setPassword] = useState('')
  const [error, setError] = useState(false)

  // Already logged in? Go home
  if (isAuthenticated()) {
    navigate({ to: '/' })
    return null
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    if (await login(password)) {
      navigate({ to: '/' })
    } else {
      setError(true)
      setPassword('')
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-stone-50">
      <div className="bg-white rounded-2xl shadow-lg p-10 w-full max-w-sm text-center">
        <h1 className="font-serif text-3xl text-stone-800 mb-2">Nicky & Sarah</h1>
        <p className="text-stone-500 text-sm mb-8">May 15, 2027</p>

        <form onSubmit={handleSubmit} className="space-y-4">
          <input
            type="password"
            placeholder="Enter password"
            value={password}
            onChange={e => { setPassword(e.target.value); setError(false) }}
            className={`w-full border rounded-lg px-4 py-2 text-stone-800 outline-none focus:ring-2 focus:ring-stone-400 ${
              error ? 'border-red-400' : 'border-stone-300'
            }`}
          />
          {error && (
            <p className="text-red-500 text-sm">Incorrect password, please try again.</p>
          )}
          <button
            type="submit"
            className="w-full bg-stone-800 text-white rounded-lg py-2 hover:bg-stone-700 transition"
          >
            Enter
          </button>
        </form>
      </div>
    </div>
  )
}