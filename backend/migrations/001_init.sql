-- AI Writer Database Schema

CREATE DATABASE IF NOT EXISTS aiwriter DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE aiwriter;

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL COMMENT '用户名',
    password VARCHAR(255) NOT NULL COMMENT '密码哈希',
    email VARCHAR(255) NOT NULL COMMENT '邮箱',
    avatar VARCHAR(500) NULL COMMENT '头像URL',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    UNIQUE KEY uk_username (username),
    UNIQUE KEY uk_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Work categories table
CREATE TABLE IF NOT EXISTS work_categories (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL COMMENT '类别名称',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Works table
CREATE TABLE IF NOT EXISTS works (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id BIGINT UNSIGNED NOT NULL COMMENT '作者ID',
    category_id BIGINT UNSIGNED NULL COMMENT '作品类别ID',
    title VARCHAR(255) NOT NULL COMMENT '书名',
    cover VARCHAR(500) NULL COMMENT '封面图URL',
    intro TEXT NULL COMMENT '简介',
    chapter_count INT UNSIGNED DEFAULT 0 COMMENT '章节数',
    word_count INT UNSIGNED DEFAULT 0 COMMENT '总字数',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    INDEX idx_user_id (user_id),
    INDEX idx_category_id (category_id),
    INDEX idx_updated_at (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Volumes table
CREATE TABLE IF NOT EXISTS volumes (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    work_id BIGINT UNSIGNED NOT NULL COMMENT '作品ID',
    volume_index INT UNSIGNED NOT NULL COMMENT '卷序号',
    title VARCHAR(255) NOT NULL COMMENT '卷名',
    chapter_count INT UNSIGNED DEFAULT 0 COMMENT '章节数',
    word_count INT UNSIGNED DEFAULT 0 COMMENT '字数',
    summary TEXT NULL COMMENT '主要情节概述',
    characters JSON NULL COMMENT '出场人物',
    plot_units JSON NULL COMMENT '主要情节单元',
    relationships TEXT NULL COMMENT '关键人物关系',
    goals TEXT NULL COMMENT '主角目标与进展',
    conflicts TEXT NULL COMMENT '冲突与转折',
    hooks TEXT NULL COMMENT '伏笔与后续钩子',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    INDEX idx_work_id (work_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Chapters table
CREATE TABLE IF NOT EXISTS chapters (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    work_id BIGINT UNSIGNED NOT NULL COMMENT '作品ID',
    volume_id BIGINT UNSIGNED NULL COMMENT '分卷ID',
    chapter_index INT UNSIGNED NOT NULL COMMENT '章序号',
    title VARCHAR(255) NOT NULL COMMENT '章节名',
    summary TEXT NULL COMMENT '剧情梗概',
    time_space VARCHAR(255) NULL COMMENT '时空设定',
    characters JSON NULL COMMENT '出场人物',
    scenes JSON NULL COMMENT '场景列表',
    word_count INT UNSIGNED DEFAULT 0 COMMENT '字数',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    INDEX idx_work_id (work_id),
    INDEX idx_volume_id (volume_id),
    INDEX idx_chapter_index (chapter_index)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Scenes table
CREATE TABLE IF NOT EXISTS scenes (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    chapter_id BIGINT UNSIGNED NOT NULL COMMENT '章节ID',
    begin TEXT NULL COMMENT '开端',
    conflict TEXT NULL COMMENT '发展与冲突',
    turn TEXT NULL COMMENT '转折',
    result TEXT NULL COMMENT '结果',
    style VARCHAR(50) NULL COMMENT '风格',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_chapter_id (chapter_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Optimization steps table
CREATE TABLE IF NOT EXISTS optimization_steps (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    name VARCHAR(100) NOT NULL COMMENT '步骤名称',
    review_prompt TEXT NULL COMMENT '审阅提示词模板',
    optimize_prompt TEXT NULL COMMENT '原文优化提示模板',
    step_order INT UNSIGNED NOT NULL COMMENT '步骤顺序',
    is_default TINYINT(1) DEFAULT 0 COMMENT '是否默认步骤',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_user_id (user_id),
    INDEX idx_step_order (step_order)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Optimization records table
CREATE TABLE IF NOT EXISTS optimization_records (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    work_id BIGINT UNSIGNED NOT NULL COMMENT '作品ID',
    chapter_id BIGINT UNSIGNED NULL COMMENT '章节ID',
    step_id BIGINT UNSIGNED NOT NULL COMMENT '步骤ID',
    original_text TEXT NULL COMMENT '原文内容',
    optimized_text TEXT NULL COMMENT '优化后内容',
    review_conclusion TEXT NULL COMMENT '审阅结论',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_work_id (work_id),
    INDEX idx_chapter_id (chapter_id),
    INDEX idx_step_id (step_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Publish tasks table
CREATE TABLE IF NOT EXISTS publish_tasks (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    work_id BIGINT UNSIGNED NOT NULL COMMENT '作品ID',
    platform VARCHAR(50) NOT NULL COMMENT '目标平台',
    chapter_ids JSON NULL COMMENT '发布章节ID列表',
    split_word_count INT UNSIGNED DEFAULT 0 COMMENT '重划分字数',
    new_chapter_names JSON NULL COMMENT '新章节名列表',
    status TINYINT UNSIGNED DEFAULT 0 COMMENT '状态：0-进行中，1-成功，2-失败',
    result TEXT NULL COMMENT '结果信息',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_work_id (work_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Notifications table
CREATE TABLE IF NOT EXISTS notifications (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    content TEXT NOT NULL COMMENT '通知内容',
    is_read TINYINT(1) DEFAULT 0 COMMENT '是否已读',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_user_id (user_id),
    INDEX idx_is_read (is_read)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert default work categories
INSERT INTO work_categories (name) VALUES 
('玄幻'), ('都市'), ('历史'), ('科幻'), ('游戏'), ('其他');
