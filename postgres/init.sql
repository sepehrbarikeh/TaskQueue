-- PostgreSQL Initialization Script for TaskQueue Engine

-- Create database if not exists
SELECT 'CREATE DATABASE taskqueue'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'taskqueue')\gexec

-- Connect to taskqueue database
\c taskqueue;

-- Create job_logs table
CREATE TABLE IF NOT EXISTS job_logs (
    id VARCHAR(255) PRIMARY KEY,
    job_id VARCHAR(255) NOT NULL,
    queue VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    payload TEXT,
    retry_count INTEGER DEFAULT 0,
    error TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_job_logs_job_id ON job_logs(job_id);
CREATE INDEX IF NOT EXISTS idx_job_logs_queue ON job_logs(queue);
CREATE INDEX IF NOT EXISTS idx_job_logs_status ON job_logs(status);
CREATE INDEX IF NOT EXISTS idx_job_logs_created_at ON job_logs(created_at);

-- Create a view for failed jobs
CREATE OR REPLACE VIEW failed_jobs AS
SELECT 
    job_id,
    queue,
    payload,
    retry_count,
    error,
    created_at
FROM job_logs 
WHERE status = 'failed'
ORDER BY created_at DESC;

-- Create a view for successful jobs
CREATE OR REPLACE VIEW successful_jobs AS
SELECT 
    job_id,
    queue,
    payload,
    retry_count,
    created_at
FROM job_logs 
WHERE status = 'success'
ORDER BY created_at DESC;

-- Grant permissions
GRANT ALL PRIVILEGES ON DATABASE taskqueue TO postgres;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO postgres;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO postgres;

-- Create a function to clean old logs
CREATE OR REPLACE FUNCTION clean_old_job_logs(days_to_keep INTEGER DEFAULT 30)
RETURNS INTEGER AS $$
DECLARE
    deleted_count INTEGER;
BEGIN
    DELETE FROM job_logs 
    WHERE created_at < CURRENT_TIMESTAMP - INTERVAL '1 day' * days_to_keep;
    
    GET DIAGNOSTICS deleted_count = ROW_COUNT;
    RETURN deleted_count;
END;
$$ LANGUAGE plpgsql;
