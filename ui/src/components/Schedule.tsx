const events = [
  {
    time: '3:00 PM',
    title: 'Ceremony',
    description: 'Join us as we exchange vows.',
  },
  {
    time: '4:00 PM',
    title: 'Cocktail Hour',
    description: 'Celebrate with drinks and light bites.',
  },
  {
    time: '5:30 PM',
    title: 'Reception Dinner',
    description: 'Dinner, toasts, and dancing.',
  },
  {
    time: '9:00 PM',
    title: 'Dancing & Celebration',
    description: 'Keep the party going on the dance floor.',
  },
  {
    time: '11:00 PM',
    title: 'Send-Off',
    description: 'Bid us farewell with sparklers.',
  },
]

export default function Schedule() {
  return (
    <section className="mx-auto max-w-2xl px-4 py-12">
      <h2 className="mb-8 text-center font-serif text-3xl font-semibold text-stone-800">
        Schedule
      </h2>
      <p className="mb-8 text-center text-stone-500 italic">
        Saturday, May 15, 2027
      </p>
      <ol className="relative border-l border-stone-200 pl-8 space-y-10">
        {events.map(({ time, title, description }) => (
          <li key={title} className="relative">
            <span className="absolute -left-[2.125rem] flex h-4 w-4 items-center justify-center rounded-full border-2 border-stone-300 bg-white ring-4 ring-white" />
            <time className="mb-1 block text-sm font-medium text-stone-400">
              {time}
            </time>
            <h3 className="text-lg font-semibold text-stone-800">{title}</h3>
            <p className="text-sm text-stone-500">{description}</p>
          </li>
        ))}
      </ol>
    </section>
  )
}
