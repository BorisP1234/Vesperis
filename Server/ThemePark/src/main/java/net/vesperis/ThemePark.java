package net.vesperis;

import net.minestom.server.MinecraftServer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class ThemePark {
    public static void main(String[] args) {
        Logger logger = LoggerFactory.getLogger(ThemePark.class);
        logger.info("Starting Vesperis ThemePark...");
        MinecraftServer.init();
        InstanceManager.loadMainInstance();
    }
}