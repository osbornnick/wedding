import { useState } from 'react'

export default function PhotoCarousel() {
  const [photoSrc, setPhotoSrc] = useState('/images/engagement.jpg')
  const [photoDescription, setPhotoDescription] = useState(
    'Dec 20 2025, Sarah and Nicky returned to Longfellow park, the site of their first date 7 years prior, and got engaged!',
  )
  return (
    <div className="flex flex-col m-5 items-center">
      <div className="flex items-center justify-items-center">
        <p>{'<'}</p>
        <img
          src={photoSrc}
          alt="Nicky and Sarah"
          className="w-auto max-h-128 sm:max-h-160 rounded-xl"
        />
        <p>{'>'}</p>
      </div>
      <p className="mt-2 text-center text-sm text-stone-500">
        {photoDescription}
      </p>
    </div>
  )
}
