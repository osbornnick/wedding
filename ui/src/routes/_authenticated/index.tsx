import { createFileRoute } from '@tanstack/react-router'
import { useEffect, useState } from 'react'
import Schedule from '../../components/Schedule'
import Travel from '../../components/Travel'
import RSVP from '../../components/RSVP'
import PhotoCarousel from '../../components/PhotoCarousel'

export const Route = createFileRoute('/_authenticated/')({ component: Home })

const WEDDING_DATE = new Date('2027-05-15T15:00:00')

function useCountdown(target: Date) {
  const calc = () => {
    const diff = target.getTime() - Date.now()
    if (diff <= 0) return { days: 0, hours: 0, minutes: 0, seconds: 0 }
    const days = Math.floor(diff / 86_400_000)
    const hours = Math.floor((diff % 86_400_000) / 3_600_000)
    const minutes = Math.floor((diff % 3_600_000) / 60_000)
    const seconds = Math.floor((diff % 60_000) / 1_000)
    return { days, hours, minutes, seconds }
  }

  const [time, setTime] = useState<ReturnType<typeof calc> | null>(null)

  useEffect(() => {
    setTime(calc())
    const id = setInterval(() => setTime(calc()), 1000)
    return () => clearInterval(id)
  }, [])

  return time
}

function CountdownUnit({ value, label }: { value: number; label: string }) {
  return (
    <div className="flex flex-col items-center">
      <span className="tabular-nums text-4xl font-bold text-stone-800 sm:text-5xl">
        {String(value).padStart(2, '0')}
      </span>
      <span className="mt-1 text-xs font-medium uppercase tracking-widest text-stone-400">
        {label}
      </span>
    </div>
  )
}

function Home() {
  const time = useCountdown(WEDDING_DATE)

  return (
    <main>
      {/* Banner */}
      <PhotoCarousel />

      {/* Date & Countdown */}
      <div className="border-b border-stone-100 bg-white px-4 py-8 text-center">
        <p className="mb-6 font-serif text-xl text-stone-500 italic">
          May 15, 2027
        </p>
        {time && (
          <div className="flex justify-center gap-8 sm:gap-12">
            <CountdownUnit value={time.days} label="Days" />
            <CountdownUnit value={time.hours} label="Hours" />
            <CountdownUnit value={time.minutes} label="Minutes" />
            <CountdownUnit value={time.seconds} label="Seconds" />
          </div>
        )}
      </div>

      {/* Full page sections */}
      <div className="divide-y divide-stone-100">
        <Schedule />
        <RSVP />
        <Travel />
      </div>
    </main>
  )
}
