import { Box } from '@chakra-ui/react'
import React from 'react'

interface VideoEmbedProps {
  embedId: string
}

export const VideoEmbed = (props: VideoEmbedProps) => {

  const { embedId } = props

  return (
    <Box  mb="5">
      <iframe
        width="500"
        height="300"
        src={`https://www.youtube.com/embed/${embedId}`}
        frameBorder="0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
        allowFullScreen
        title="Embedded youtube"
      />
    </Box>
  )
}