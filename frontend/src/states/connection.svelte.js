export let connection = $state({
  current_connection: {
    connection_id: "",
    connection_name: "",
  },
  active_connections: [],
  credential: {
    connection_name: "",
    db_vendor: "PostgreSQL",
    hex_color: "",
    host: "",
    port: "",
    user: "",
    password: "",
    database_name: "",
  },
});
