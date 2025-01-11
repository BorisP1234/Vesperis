package net.vesperis;

import net.minestom.server.MinecraftServer;
import net.minestom.server.entity.Player;
import net.minestom.server.instance.InstanceContainer;
import net.minestom.server.instance.SharedInstance;
import net.minestom.server.instance.anvil.AnvilLoader;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class InstanceManager {

    private final InstanceContainer mainInstance;
    private final Logger logger;

    public InstanceManager() {
        logger = LoggerFactory.getLogger(InstanceManager.class);

        try {
            net.minestom.server.instance.InstanceManager manager = MinecraftServer.getInstanceManager();
            mainInstance = manager.createInstanceContainer();
            mainInstance.setChunkLoader(new AnvilLoader("worlds/themepark/main"));

            logger.info("Successfully loaded main instance.");
            throw new Exception("HaHAHA");


        } catch (Exception e) {
            logger.error("Error loading main instance: {}", e, e);
            throw new RuntimeException(e);
        }

    }

    public static void loadMainInstance() {
        InstanceManager instanceManager = new InstanceManager();
    }

    public void loadInstance(Player player) {
        SharedInstance instance = new SharedInstance(player.getUuid(), mainInstance);
    }

}
