export interface RegistryItem {
  name: string
  url: string
  description: string
}

export const defaultRegistries: RegistryItem[] = [
  {
    name: 'Crate & Barrel',
    url: 'https://www.crateandbarrel.com',
    description: 'Kitchen, dining, and home essentials.',
  },
  {
    name: 'Williams Sonoma',
    url: 'https://www.williams-sonoma.com',
    description: 'Cookware, bakeware, and entertaining.',
  },
  {
    name: 'Zola',
    url: 'https://www.zola.com',
    description: 'Our full gift registry and honeymoon fund.',
  },
]

export default function Registry({
  items = defaultRegistries,
}: {
  items?: RegistryItem[]
}) {
  return (
    <section className="mx-auto max-w-2xl px-4 py-12">
      <h2 className="mb-4 text-center font-serif text-3xl font-semibold text-stone-800">
        Registry
      </h2>
      <p className="mb-10 text-center text-stone-500">
        Your presence is the greatest gift. If you'd like to celebrate with a
        gift, we're registered at the places below.
      </p>
      <div className="space-y-4">
        {items.map(({ name, url, description }) => (
          <a
            key={name}
            href={url}
            target="_blank"
            rel="noopener noreferrer"
            className="flex items-center justify-between rounded-2xl border border-stone-200 bg-stone-50 p-6 no-underline transition hover:border-stone-300 hover:shadow-sm"
          >
            <div>
              <h3 className="text-base font-semibold text-stone-800">{name}</h3>
              <p className="mt-0.5 text-sm text-stone-500">{description}</p>
            </div>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
              className="h-5 w-5 flex-shrink-0 text-stone-400"
            >
              <path
                fillRule="evenodd"
                d="M5.22 14.78a.75.75 0 001.06 0l7.22-7.22v5.69a.75.75 0 001.5 0v-7.5a.75.75 0 00-.75-.75h-7.5a.75.75 0 000 1.5h5.69l-7.22 7.22a.75.75 0 000 1.06z"
                clipRule="evenodd"
              />
            </svg>
          </a>
        ))}
      </div>
    </section>
  )
}
