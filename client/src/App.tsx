import React from 'react';

import { ChakraProvider, Flex, Box } from '@chakra-ui/react'
import { Chat } from './components/Chat.tsx';
import { VideoEmbed } from './components/VideoEmbed.tsx';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

function App() {

  return (
    <ChakraProvider>
      <Flex justify="center">
        <Box maxW="lg" w="full">
          <VideoEmbed embedId='bBjENsfyvFs' />
          <Chat />
        </Box>
      </Flex>
      <ToastContainer />
    </ChakraProvider>

  );
}

export default App;
