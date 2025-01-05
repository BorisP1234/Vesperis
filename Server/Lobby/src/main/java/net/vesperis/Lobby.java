package net.vesperis;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Lobby {

    private static Lobby lobby;
    private final Logger logger;

    private Lobby() {
        logger = LoggerFactory.getLogger(Lobby.class);
    }

    public static void main(String[] args) {
        lobby = new Lobby();
        lobby.logger.info("Starting Vesperis Lobby...");
    }
}