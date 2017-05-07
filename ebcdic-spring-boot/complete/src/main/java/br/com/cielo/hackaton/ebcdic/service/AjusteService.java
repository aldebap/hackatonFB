package br.com.cielo.hackaton.ebcdic.service;

import java.sql.SQLException;

import br.com.cielo.hackaton.ebcdic.domain.SettlementMovement;

public interface AjusteService {
	
	public void createMovement(SettlementMovement movement) throws SQLException;

}
