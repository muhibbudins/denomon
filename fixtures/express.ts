import * as expressive from "https://raw.githubusercontent.com/NMathar/deno-express/master/mod.ts";

(async () => {
  const port = 3000;
  const app = new expressive.App();

  app.use(expressive.simpleLog());
  app.use(expressive.bodyParser.json());

  app.get('/', async (req, res) => {
    res.json({foo: "baz"})
  })

  const server = await app.listen(port);

  console.log("app listening on port " + server.port);
})();