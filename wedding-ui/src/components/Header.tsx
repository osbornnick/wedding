import { Link, useNavigate } from '@tanstack/react-router'
import { logout } from '../lib/auth'

const navLinks = [
  { to: '/', label: 'Home', exact: true },
  { to: '/schedule', label: 'Schedule' },
  { to: '/travel', label: 'Travel' },
  { to: '/rsvp', label: 'RSVP' },
  { to: '/registry', label: 'Registry' },
] as const

export default function Header() {
  const navigate = useNavigate()

  function handleLogout() {
    logout()
    navigate({ to: '/login' })
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
            <button
              onClick={handleLogout}
              className="rounded-md bg-stone-100 px-3 py-1 text-xs text-stone-500 transition hover:bg-stone-200 hover:text-stone-800"
            >
              Dev: Logout
            </button>
          )}
        </div>
      </nav>
    </header>
  )
}
