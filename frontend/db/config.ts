import { column, defineDb, defineTable } from "astro:db";

const Test = defineTable({
	columns: {
		text: column.text(),
	},
});

// https://astro.build/db/config
export default defineDb({
	tables: {},
});
