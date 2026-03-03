import { createFileRoute, redirect, useNavigate } from '@tanstack/react-router'
import { useState } from 'react'
import { loginFn } from '../lib/auth'
import { Button, Card, HelperText, TextInput } from 'flowbite-react'

export const Route = createFileRoute('/login')({
  beforeLoad: async ({ context }) => {
    if (context.isAuthenticated) {
      throw redirect({ to: '/' })
    }
  },
  component: LoginPage,
})

function LoginPage() {
  const navigate = useNavigate()
  const [password, setPassword] = useState('')
  const [error, setError] = useState(false)

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    if (await loginFn({ data: { password } })) {
      navigate({ to: '/' })
    } else {
      setError(true)
      setPassword('')
    }
  }

  return (
    <div className="grow flex items-center justify-center">
      <Card
        className="max-w-sm"
        imgSrc="/images/engagement.jpg"
        imgAlt="Engagement photo"
      >
        <form onSubmit={handleSubmit} className="flex flex-col gap-4 p-2">
          <div className="text-center">
            <h1 className="font-serif text-3xl mb-2">Nicky & Sarah</h1>
            <p className="text-sm mb-8">May 15, 2027</p>
          </div>
          <TextInput
            type="password"
            placeholder="Enter password"
            value={password}
            onChange={(e) => {
              setPassword(e.target.value)
              setError(false)
            }}
            color={error ? 'failure' : undefined}
          />
          <HelperText className="font-medium" hidden={!error}>
            Incorrect password, please try again.
          </HelperText>
          <Button type="submit" className="w-full py-2 transition">
            Enter
          </Button>
        </form>
      </Card>
    </div>
  )
}
