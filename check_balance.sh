balance_url="https://api-arabica.celenium.io/v1/address/$ADDRESS"
faucet_url="https://faucet.celestia-arabica-11.com/api/v1/faucet/give_me"

while true; do
    response=$(curl -s $balance_url)
    spendable=$(echo "$response" | jq -r '.balance.spendable')
    if [ "$spendable" -lt "$THRESHOLD" ]; then
        echo "$(date '+%Y-%m-%d %H:%M:%S') Low balance: $spendable < $THRESHOLD. Requesting TIA from faucet."
        curl -sS -X POST $faucet_url \
            -H "Content-Type: application/json" \
            -d "{\"address\":\"$ADDRESS\",\"chainId\":\"$CHAIN_ID\"}"
    else
        echo "$(date '+%Y-%m-%d %H:%M:%S') Enough balance: $spendable > $THRESHOLD"
    fi

    sleep $INTERVAL
done
