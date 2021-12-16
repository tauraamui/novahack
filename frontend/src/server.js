import { createServer } from "miragejs";

createServer({
  routes() {
    this.namespace = "api";

    this.post("/event", (schema, request) => {
      let room = JSON.parse(request.requestBody);
      room.id = Math.floor(Math.random() * 100);

      return { event: room };
    });
  },
});
