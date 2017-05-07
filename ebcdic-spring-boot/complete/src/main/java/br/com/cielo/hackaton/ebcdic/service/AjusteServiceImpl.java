package br.com.cielo.hackaton.ebcdic.service;

import java.sql.SQLException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import br.com.cielo.hackaton.ebcdic.domain.SettlementMovement;
import br.com.cielo.hackaton.ebcdic.persistence.SettlementMovementRepository;

@Service
public class AjusteServiceImpl implements AjusteService {
	
	@Autowired
	SettlementMovementRepository repository;

	@Override
	public void createMovement(SettlementMovement movement) throws SQLException {
		repository.insertMovement(movement);
	}
}
