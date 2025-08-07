package postgres

import "log"

func (p PostgresDB) Migrate() {
	err := p.DB.AutoMigrate()
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}
	log.Println("✅ Migration completed.")
}
