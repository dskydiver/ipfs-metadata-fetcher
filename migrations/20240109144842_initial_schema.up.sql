BEGIN;

-- CreateTable
CREATE TABLE "metadata" (
    "id" TEXT NOT NULL DEFAULT ulid_generate(),
    "cid" TEXT NOT NULL,
    "image" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(3),

    CONSTRAINT "metadata_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "metadata_cid_key" ON "metadata"("cid");

COMMIT;