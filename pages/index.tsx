import type { NextPage } from 'next'
import Head from 'next/head'
import ChatWindow from '../components/ChatWindow'
import Loading from '../components/Loading'
import { useWebsocket } from '../util/context'

const Home: NextPage = () => {
  const { connection, user } = useWebsocket()

  return (
    <>
      <Head>
        <title>Chat</title>
      </Head>
      <main className="h-full">
        {!user || !connection ? <Loading /> : <ChatWindow />}
        <footer></footer>
      </main>
    </>
  )
}

export default Home
