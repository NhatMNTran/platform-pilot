import axios from "axios";
import yargs from "yargs";
import { hideBin } from "yargs/helpers";

yargs(hideBin(process.argv))
    .command(
        "env-status",
        "Get environment status",
        {
            id: {
                type: "string",
                demandOption: true,
            },
        },
        async (argv: any) => {
            const response = await axios.get(
                `http://localhost:8080/environments/${argv.id}`
            );

            console.log(JSON.stringify(response.data, null, 2));
        }
    )
    .help()
    .parse();