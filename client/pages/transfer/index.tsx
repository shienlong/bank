import Link from 'next/link'
import fetch from 'isomorphic-unfetch'
import { NextPage } from 'next';

type Props = {
    uuid: string
    currency: string
    amount: number
}

const Index: NextPage<Props> = ({ uuid, currency, amount }) => (
    <div>
        <h1>Balance</h1>
        <p>Account ID: {uuid}</p>
        <p>Amount: {currency} {amount}</p>
    </div>
)

Index.getInitialProps = async () => {
    const res = await fetch(process.env.BALANCE_SERVICE_URL)
    const data : Props = await res.json()

    return { ...data }
}

export default Index