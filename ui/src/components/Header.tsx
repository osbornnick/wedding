import { Link, useNavigate } from '@tanstack/react-router'
import { DarkThemeToggle, Button } from 'flowbite-react'

const navLinks = [
  { to: '/', label: 'Home', exact: true },
  { to: '/schedule', label: 'Schedule' },
  { to: '/travel', label: 'Travel' },
  { to: '/rsvp', label: 'RSVP' },
  { to: '/registry', label: 'Registry' },
] as const

export default function Header() {
  const navigate = useNavigate()

  async function handleLogout() {
    navigate({ to: '/logout' })
  }

  return (
    <header className="sticky top-0 z-50 border-b border-stone-200 bg-white/90 px-4 backdrop-blur-md shadow-sm">
      <nav className="mx-auto flex max-w-5xl items-center justify-between py-4">
        <Link
          to="/"
          className="font-serif text-xl font-semibold tracking-wide text-stone-800 no-underline"
        >
          Nicky &amp; Sarah
        </Link>

        <div className="flex items-center gap-6 text-sm font-medium">
          {navLinks.map(({ to, label, exact }) => (
            <Link
              key={to}
              to={to}
              className="text-stone-500 transition hover:text-stone-900"
              activeProps={{ className: 'text-stone-900 font-semibold' }}
              activeOptions={exact ? { exact: true } : undefined}
            >
              {label}
            </Link>
          ))}
          {import.meta.env.DEV && (
            <Button onClick={handleLogout}>Dev: Logout</Button>
          )}
          <DarkThemeToggle />
        </div>
      </nav>
    </header>
  )
}
